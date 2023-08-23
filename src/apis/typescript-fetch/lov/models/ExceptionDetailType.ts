/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud List of Values Management API
 * APIs to cater for List of Value functionality in OPERA Cloud. A List of Values in the OPERA Application can be configured by a property.  Then by using these APIs you can retrieve all configured codes.  As an example, Titles is a configurable ListOfValues.  A hotel can specify what titles they wish to use, and thus fetching the LOV for title, you can view the codes that are configured for a property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ErrorInstance } from './ErrorInstance';
import {
    ErrorInstanceFromJSON,
    ErrorInstanceFromJSONTyped,
    ErrorInstanceToJSON,
} from './ErrorInstance';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';

/**
 * Complex type that contains error details for a REST call.
 * @export
 * @interface ExceptionDetailType
 */
export interface ExceptionDetailType {
    /**
     * Absolute URI [RFC3986] that identifies the problem type.  When dereferenced, it SHOULD provide a human-readable summary of the problem (for example, using HTML).
     * @type {string}
     * @memberof ExceptionDetailType
     */
    type: string;
    /**
     * Short, human-readable summary of the problem.  The summary SHOULD NOT change for subsequent occurrences of the problem, except for purposes of localization.
     * @type {string}
     * @memberof ExceptionDetailType
     */
    title: string;
    /**
     * HTTP status code for this occurrence of the problem, set by the origin server.
     * @type {number}
     * @memberof ExceptionDetailType
     */
    status?: number;
    /**
     * Human-readable description specific to this occurrence of the problem.
     * @type {string}
     * @memberof ExceptionDetailType
     */
    detail?: string;
    /**
     * Absolute URI that identifies the specific occurrence of the problem.  It may or may not provide additional information if dereferenced.
     * @type {string}
     * @memberof ExceptionDetailType
     */
    instance?: string;
    /**
     * Application error code, which is different from HTTP error code.
     * @type {string}
     * @memberof ExceptionDetailType
     */
    oerrorCode?: string;
    /**
     * Path to the problem at the resource or property level.
     * @type {string}
     * @memberof ExceptionDetailType
     */
    oerrorPath?: string;
    /**
     * Details of the error message, consisting of a hierarchical tree structure.
     * @type {Array<ErrorInstance>}
     * @memberof ExceptionDetailType
     */
    oerrorDetails?: Array<ErrorInstance>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ExceptionDetailType
     */
    links?: Array<InstanceLink>;
}

/**
 * Check if a given object implements the ExceptionDetailType interface.
 */
export function instanceOfExceptionDetailType(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "type" in value;
    isInstance = isInstance && "title" in value;

    return isInstance;
}

export function ExceptionDetailTypeFromJSON(json: any): ExceptionDetailType {
    return ExceptionDetailTypeFromJSONTyped(json, false);
}

export function ExceptionDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExceptionDetailType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'type': json['type'],
        'title': json['title'],
        'status': !exists(json, 'status') ? undefined : json['status'],
        'detail': !exists(json, 'detail') ? undefined : json['detail'],
        'instance': !exists(json, 'instance') ? undefined : json['instance'],
        'oerrorCode': !exists(json, 'o:errorCode') ? undefined : json['o:errorCode'],
        'oerrorPath': !exists(json, 'o:errorPath') ? undefined : json['o:errorPath'],
        'oerrorDetails': !exists(json, 'o:errorDetails') ? undefined : ((json['o:errorDetails'] as Array<any>).map(ErrorInstanceFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
    };
}

export function ExceptionDetailTypeToJSON(value?: ExceptionDetailType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'type': value.type,
        'title': value.title,
        'status': value.status,
        'detail': value.detail,
        'instance': value.instance,
        'o:errorCode': value.oerrorCode,
        'o:errorPath': value.oerrorPath,
        'o:errorDetails': value.oerrorDetails === undefined ? undefined : ((value.oerrorDetails as Array<any>).map(ErrorInstanceToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
    };
}

