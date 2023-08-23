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
import type { ActionStatusType } from './ActionStatusType';
import {
    ActionStatusTypeFromJSON,
    ActionStatusTypeFromJSONTyped,
    ActionStatusTypeToJSON,
} from './ActionStatusType';
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';

/**
 * Identifies the response expected from staff or other parties relative to a Track It ticket.
 * @export
 * @interface TrackItActionType
 */
export interface TrackItActionType {
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof TrackItActionType
     */
    trackItAction?: CodeDescriptionType;
    /**
     * 
     * @type {ActionStatusType}
     * @memberof TrackItActionType
     */
    status?: ActionStatusType;
}

/**
 * Check if a given object implements the TrackItActionType interface.
 */
export function instanceOfTrackItActionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TrackItActionTypeFromJSON(json: any): TrackItActionType {
    return TrackItActionTypeFromJSONTyped(json, false);
}

export function TrackItActionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TrackItActionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'trackItAction': !exists(json, 'trackItAction') ? undefined : CodeDescriptionTypeFromJSON(json['trackItAction']),
        'status': !exists(json, 'status') ? undefined : ActionStatusTypeFromJSON(json['status']),
    };
}

export function TrackItActionTypeToJSON(value?: TrackItActionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'trackItAction': CodeDescriptionTypeToJSON(value.trackItAction),
        'status': ActionStatusTypeToJSON(value.status),
    };
}

