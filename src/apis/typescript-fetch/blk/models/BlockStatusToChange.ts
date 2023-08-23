/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ChangeBlockStatusType } from './ChangeBlockStatusType';
import {
    ChangeBlockStatusTypeFromJSON,
    ChangeBlockStatusTypeFromJSONTyped,
    ChangeBlockStatusTypeToJSON,
} from './ChangeBlockStatusType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Request object for changing the booking status of the business block.
 * @export
 * @interface BlockStatusToChange
 */
export interface BlockStatusToChange {
    /**
     * Contains the booking status of the business block.
     * @type {object}
     * @memberof BlockStatusToChange
     */
    blocksStatus?: object;
    /**
     * 
     * @type {ChangeBlockStatusType}
     * @memberof BlockStatusToChange
     */
    changeBlockStatus?: ChangeBlockStatusType;
    /**
     * Indicator if the request is a verification on whether the block status can be changed.
     * @type {boolean}
     * @memberof BlockStatusToChange
     */
    verificationOnly?: boolean;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof BlockStatusToChange
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof BlockStatusToChange
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the BlockStatusToChange interface.
 */
export function instanceOfBlockStatusToChange(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockStatusToChangeFromJSON(json: any): BlockStatusToChange {
    return BlockStatusToChangeFromJSONTyped(json, false);
}

export function BlockStatusToChangeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockStatusToChange {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'blocksStatus': !exists(json, 'blocksStatus') ? undefined : json['blocksStatus'],
        'changeBlockStatus': !exists(json, 'changeBlockStatus') ? undefined : ChangeBlockStatusTypeFromJSON(json['changeBlockStatus']),
        'verificationOnly': !exists(json, 'verificationOnly') ? undefined : json['verificationOnly'],
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function BlockStatusToChangeToJSON(value?: BlockStatusToChange | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'blocksStatus': value.blocksStatus,
        'changeBlockStatus': ChangeBlockStatusTypeToJSON(value.changeBlockStatus),
        'verificationOnly': value.verificationOnly,
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

