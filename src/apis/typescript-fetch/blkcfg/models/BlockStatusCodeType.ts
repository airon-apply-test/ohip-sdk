/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block Configuration API
 * APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { BlockStatusCodeStatusType } from './BlockStatusCodeStatusType';
import {
    BlockStatusCodeStatusTypeFromJSON,
    BlockStatusCodeStatusTypeFromJSONTyped,
    BlockStatusCodeStatusTypeToJSON,
} from './BlockStatusCodeStatusType';
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';
import type { StatusColorType } from './StatusColorType';
import {
    StatusColorTypeFromJSON,
    StatusColorTypeFromJSONTyped,
    StatusColorTypeToJSON,
} from './StatusColorType';

/**
 * Block status code information.
 * @export
 * @interface BlockStatusCodeType
 */
export interface BlockStatusCodeType {
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof BlockStatusCodeType
     */
    status?: CodeDescriptionType;
    /**
     * 
     * @type {BlockStatusCodeStatusType}
     * @memberof BlockStatusCodeType
     */
    roomStatusType?: BlockStatusCodeStatusType;
    /**
     * 
     * @type {BlockStatusCodeStatusType}
     * @memberof BlockStatusCodeType
     */
    cateringStatusType?: BlockStatusCodeStatusType;
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof BlockStatusCodeType
     */
    defaultReservationType?: CodeDescriptionType;
    /**
     * Reason type for every block reservation cancellation. Block Cancellation Reason, Block Refused Reasons, and Block Lost Reasons.
     * @type {string}
     * @memberof BlockStatusCodeType
     */
    reasonType?: string;
    /**
     * 
     * @type {StatusColorType}
     * @memberof BlockStatusCodeType
     */
    color?: StatusColorType;
    /**
     * Indicates if the rooms with this status code will be allowed for pickup or not.
     * @type {boolean}
     * @memberof BlockStatusCodeType
     */
    allowPickup?: boolean;
    /**
     * Indicates if the rooms with this status code can be returned to availability or not.
     * @type {boolean}
     * @memberof BlockStatusCodeType
     */
    returnToInventory?: boolean;
    /**
     * Indicates if the status code is the starting status of the status cycle or not.
     * @type {boolean}
     * @memberof BlockStatusCodeType
     */
    starting?: boolean;
    /**
     * Indicates the default status for all new leads. Only one lead status can be selected as the default.
     * @type {boolean}
     * @memberof BlockStatusCodeType
     */
    leadStatus?: boolean;
    /**
     * Indicates if the status code logs the catering changes or not.
     * @type {boolean}
     * @memberof BlockStatusCodeType
     */
    logCatering?: boolean;
    /**
     * Indicates if the function diary will be shown within the Sales and Catering.
     * @type {boolean}
     * @memberof BlockStatusCodeType
     */
    showDiary?: boolean;
    /**
     * Indicates the order the status code is displayed in the status cycle.
     * @type {number}
     * @memberof BlockStatusCodeType
     */
    orderBy?: number;
    /**
     * Returns true if the status code is already used in Block / Event / Status Flow.
     * @type {boolean}
     * @memberof BlockStatusCodeType
     */
    inUse?: boolean;
    /**
     * Returns true if the status code is already used in Events.
     * @type {boolean}
     * @memberof BlockStatusCodeType
     */
    cateringInUse?: boolean;
    /**
     * Indicates that the business block created with this Status is an Opportunity.
     * @type {boolean}
     * @memberof BlockStatusCodeType
     */
    opportunity?: boolean;
}

/**
 * Check if a given object implements the BlockStatusCodeType interface.
 */
export function instanceOfBlockStatusCodeType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockStatusCodeTypeFromJSON(json: any): BlockStatusCodeType {
    return BlockStatusCodeTypeFromJSONTyped(json, false);
}

export function BlockStatusCodeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockStatusCodeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'status': !exists(json, 'status') ? undefined : CodeDescriptionTypeFromJSON(json['status']),
        'roomStatusType': !exists(json, 'roomStatusType') ? undefined : BlockStatusCodeStatusTypeFromJSON(json['roomStatusType']),
        'cateringStatusType': !exists(json, 'cateringStatusType') ? undefined : BlockStatusCodeStatusTypeFromJSON(json['cateringStatusType']),
        'defaultReservationType': !exists(json, 'defaultReservationType') ? undefined : CodeDescriptionTypeFromJSON(json['defaultReservationType']),
        'reasonType': !exists(json, 'reasonType') ? undefined : json['reasonType'],
        'color': !exists(json, 'color') ? undefined : StatusColorTypeFromJSON(json['color']),
        'allowPickup': !exists(json, 'allowPickup') ? undefined : json['allowPickup'],
        'returnToInventory': !exists(json, 'returnToInventory') ? undefined : json['returnToInventory'],
        'starting': !exists(json, 'starting') ? undefined : json['starting'],
        'leadStatus': !exists(json, 'leadStatus') ? undefined : json['leadStatus'],
        'logCatering': !exists(json, 'logCatering') ? undefined : json['logCatering'],
        'showDiary': !exists(json, 'showDiary') ? undefined : json['showDiary'],
        'orderBy': !exists(json, 'orderBy') ? undefined : json['orderBy'],
        'inUse': !exists(json, 'inUse') ? undefined : json['inUse'],
        'cateringInUse': !exists(json, 'cateringInUse') ? undefined : json['cateringInUse'],
        'opportunity': !exists(json, 'opportunity') ? undefined : json['opportunity'],
    };
}

export function BlockStatusCodeTypeToJSON(value?: BlockStatusCodeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'status': CodeDescriptionTypeToJSON(value.status),
        'roomStatusType': BlockStatusCodeStatusTypeToJSON(value.roomStatusType),
        'cateringStatusType': BlockStatusCodeStatusTypeToJSON(value.cateringStatusType),
        'defaultReservationType': CodeDescriptionTypeToJSON(value.defaultReservationType),
        'reasonType': value.reasonType,
        'color': StatusColorTypeToJSON(value.color),
        'allowPickup': value.allowPickup,
        'returnToInventory': value.returnToInventory,
        'starting': value.starting,
        'leadStatus': value.leadStatus,
        'logCatering': value.logCatering,
        'showDiary': value.showDiary,
        'orderBy': value.orderBy,
        'inUse': value.inUse,
        'cateringInUse': value.cateringInUse,
        'opportunity': value.opportunity,
    };
}

